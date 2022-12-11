package service

import (
	"context"
	"time"

	"github.com/social-nyetwork/backend/internal/models"
)

func CreateMessage(data *models.CreateMessageRequest) (id string, err error) {
	err = validate.Struct(data)
	if err == nil {
		generatedId := dbpool.QueryRow(context.Background(),
			"INSERT INTO message(sender, recipient, text) VALUES($1, $2, $3) RETURNING id", data.Sender, data.Recipient, data.Text)

		err = generatedId.Scan(&id)
	}

	return id, err
}

func DeleteMessage(id string) error {
	result, err := dbpool.Exec(context.Background(),
		"UPDATE message SET exist = FALSE WHERE id = $1 AND active", id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return NotFoundError
	}
	return nil
}

func EditMessage(id string, from int64, data *models.MessageData) error {
	err := validate.Struct(data)
	if err == nil {
		result, err := dbpool.Exec(context.Background(),
			"UPDATE message SET text = $3 WHERE id = $1 AND sender = $2 AND active", id, from, data.Text)
		if err != nil {
			return err
		}
		if result.RowsAffected() == 0 {
			return NotFoundError
		}
	}
	return err
}

func GetMessages(from int64, to int64, page int, pageSize int) (response *models.GetMessagesResponse, err error) {
	result, err := dbpool.Query(context.Background(), "SELECT \"id\", created, last_modified, sender, recipient, text FROM message WHERE (sender = $1 AND recipient = $2 OR sender = $2 AND recipient = $1) AND active LIMIT $3 OFFSET $4", from, to, pageSize, page*pageSize)
	if err == nil {
		response = &models.GetMessagesResponse{Items: []*models.MessageItem{}, Page: &models.Page{PageNumber: page, PageSize: pageSize, TotalItems: 0}}
		for result.Next() {
			var id, text string
			var sender, recipient int64
			var sent, edited time.Time
			if innerErr := result.Scan(&id, &sent, &edited, &sender, &recipient, &text); innerErr == nil {
				response.Items = append(response.Items, &models.MessageItem{ID: id, Sent: sent, Edited: edited, IsReceived: false, Sender: sender, Recipient: recipient, Text: text, OriginalID: nil})
			}
		}
	}

	return response, err
}
