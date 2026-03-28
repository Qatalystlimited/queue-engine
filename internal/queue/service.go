package queue

import (
    "context"
    "database/sql"
    "fmt"
    "time"
)

type Service struct {
    DB *sql.DB
}

func (s *Service) JoinQueue(ctx context.Context, userID, queueID string) (string, int32, error) {
    tx, err := s.DB.BeginTx(ctx, nil)
    if err != nil {
        return "", 0, err
    }
    defer tx.Rollback()

    // Get next position
    var position int32
    err = tx.QueryRowContext(ctx,
        `SELECT COALESCE(MAX(position), 0) + 1 FROM queue_entries WHERE queue_id = $1`,
        queueID,
    ).Scan(&position)
    if err != nil {
        return "", 0, err
    }

    // Generate ticket ID
    ticketID := fmt.Sprintf("T-%s-%d", queueID, time.Now().UnixMilli())

    // Insert entry
    _, err = tx.ExecContext(ctx,
        `INSERT INTO queue_entries (ticket_id, user_id, queue_id, position, status)
         VALUES ($1, $2, $3, $4, 'waiting')`,
        ticketID, userID, queueID, position,
    )
    if err != nil {
        return "", 0, err
    }

    return ticketID, position, tx.Commit()
}

func (s *Service) GetPosition(ctx context.Context, userID, queueID string) (int32, error) {
    var position int32
    err := s.DB.QueryRowContext(ctx,
        `SELECT position FROM queue_entries WHERE user_id = $1 AND queue_id = $2`,
        userID, queueID,
    ).Scan(&position)
    if err == sql.ErrNoRows {
        return 0, fmt.Errorf("user not in queue")
    }
    return position, err
}
