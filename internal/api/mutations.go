package api

import (
	"fmt"
	"strings"
)

func BuildMediaEntryUpdateMutation(params map[string]any) string {
	var argsDef strings.Builder
	var argsCall strings.Builder

	argsDef.WriteString("$id: Int")
	argsCall.WriteString("mediaId: $id")

	if _, ok := params["progress"]; ok {
		argsDef.WriteString(", $progress: Int")
		argsCall.WriteString(", progress: $progress")
	}
	if _, ok := params["status"]; ok {
		argsDef.WriteString(", $status: MediaListStatus")
		argsCall.WriteString(", status: $status")
	}
	if _, ok := params["score"]; ok {
		argsDef.WriteString(", $score: Float")
		argsCall.WriteString(", score: $score")
	}
	if _, ok := params["notes"]; ok {
		argsDef.WriteString(", $notes: String")
		argsCall.WriteString(", notes: $notes")
	}

	hasCompletedAt := false
	if _, ok := params["cDate"]; ok {
		hasCompletedAt = true
	}
	if _, ok := params["cMonth"]; ok {
		hasCompletedAt = true
	}
	if _, ok := params["cYear"]; ok {
		hasCompletedAt = true
	}

	if hasCompletedAt {
		argsDef.WriteString(", $cDate: Int, $cMonth: Int, $cYear: Int")
		argsCall.WriteString(", completedAt: {day: $cDate, month: $cMonth, year: $cYear}")
	}

	hasStartedAt := false
	if _, ok := params["sDate"]; ok {
		hasStartedAt = true
	}
	if _, ok := params["sMonth"]; ok {
		hasStartedAt = true
	}
	if _, ok := params["sYear"]; ok {
		hasStartedAt = true
	}

	if hasStartedAt {
		argsDef.WriteString(", $sDate: Int, $sMonth: Int, $sYear: Int")
		argsCall.WriteString(", startedAt: {day: $sDate, month: $sMonth, year: $sYear}")
	}

	return fmt.Sprintf(`mutation(%s) {
    SaveMediaListEntry(%s) {
		media {
			id
			title {
				userPreferred
			}
		}
    }
}`, argsDef.String(), argsCall.String())
}
