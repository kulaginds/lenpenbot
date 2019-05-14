package top

import (
	"fmt"
	"strings"
	"time"

	"github.com/kulaginds/lenpenbot/pkg/penis"
	"github.com/kulaginds/lenpenbot/pkg/store"
	"github.com/kulaginds/lenpenbot/pkg/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Top struct {
	bot   types.BotAPIImplementation
	store store.Store
}

type TopImplementation interface {
	GenerateToday(chatID int64, today time.Time) (string, error)
	GenerateTop(chatID int64) (string, error)
}

func NewTop(bot types.BotAPIImplementation, store store.Store) TopImplementation {
	return &Top{bot: bot, store: store}
}

func (t *Top) extractUserIDsFromEnlarges(enlarges []*store.Enlarge) []int {
	userIDs := make([]int, 0)

	for _, enlarge := range enlarges {
		userIDs = append(userIDs, enlarge.UserID)
	}

	return userIDs
}

func (t *Top) getNicksMap(userIDs []int, chatID int64) (map[int]string, error) {
	m := map[int]string{}

	for _, userID := range userIDs {
		member, err := t.bot.GetChatMember(tgbotapi.ChatConfigWithUser{
			ChatID: chatID,
			UserID: userID,
		})
		if err != nil {
			return nil, err
		}

		m[userID] = member.User.UserName
	}

	return m, nil
}

func (t *Top) generateTop(enlarges []*store.Enlarge, nicksMap map[int]string) string {
	var sb strings.Builder

	for i, enlarge := range enlarges {
		sb.Write([]byte(fmt.Sprintf("%d. %s %d cm\n", i+1, nicksMap[enlarge.UserID], enlarge.Length)))

		if i == 0 {
			sb.Write([]byte(penis.Generate(enlarge.Length)))
			sb.Write([]byte("\n---------------------------\n"))
		}
	}

	return strings.TrimRight(sb.String(), "\n")
}
