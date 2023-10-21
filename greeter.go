package eng

import (
"github.com/rs/zerolog"
"github.com/rs/zerolog/log"
)

type greeting string

func (g greeting) Greet() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

    log.Error().Msg("Error message")
    log.Warn().Msg("Warning message")
    log.Info().Msg("Info message")
    log.Debug().Msg("ENGLADDD Debug message")
    log.Trace().Msg("Trace message")

}

// exported
var Greeter greeting
