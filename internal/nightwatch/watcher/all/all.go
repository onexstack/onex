package all

//nolint: golint
import (
	_ "github.com/onexstack/onex/internal/nightwatch/watcher/cronjob/cronjob"
	_ "github.com/onexstack/onex/internal/nightwatch/watcher/cronjob/statesync"
	_ "github.com/onexstack/onex/internal/nightwatch/watcher/job/llmtrain"
	_ "github.com/onexstack/onex/internal/nightwatch/watcher/secretsclean"
	_ "github.com/onexstack/onex/internal/nightwatch/watcher/user"
)
