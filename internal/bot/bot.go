package bot

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

var db Datastore

func doesCodeExist(code string) (bool, error){
	if c, err := db.FindCodeByCode(code); err != nil {
		return false, err
	} else if *c != (Code{}) {
		return true, nil
	}

	return false, nil
}

// TODO :: Call matt1484/bl3_auto_vip to check validity
func isCodeValid() bool {
	return true
}

func AddCode(codetype string, code string)  {
	if codeExists, err := doesCodeExist(code); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error while check if code exists: %s"))
		return
	} else if codeExists {
		log.Warn().Msg("Code exists")
		return
	}

	if !isCodeValid() {
		log.Warn().Msg("Code not valid")
		return
	}

	if err := db.InsertCode(code); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error while check if code exists: %s"))
		return
	}

	log.Info().Msg(fmt.Sprintf("%s added", code))
}