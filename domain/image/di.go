package image

import (
	"wsw/backend/app/config"
	"wsw/backend/app/di"
	"wsw/backend/domain/image/path"
	"wsw/backend/domain/image/process"
	processPathProvider "wsw/backend/domain/image/process/path"
	"wsw/backend/domain/image/process/processor"
	"wsw/backend/domain/image/process/runner"
	"wsw/backend/domain/image/process/runner/result"
	"wsw/backend/domain/image/storage"
	"wsw/backend/domain/image/url"
	imageModel "wsw/backend/model/image"
)

func InitImageModuleDI(config config.Config) {
	di.InitModule("image",

		func() path.FilenameGenerator { return path.NewFilenameProvider() },
		func() path.PathProvider { return path.NewPathProvider(config.App.UploadPath) },

		func(filenameGenerator path.FilenameGenerator, pathProvider path.PathProvider) storage.Storage {
			return storage.NewUploadStorage(filenameGenerator, pathProvider)
		},
		func(pathProvider path.PathProvider, pathGenerator path.FilenameGenerator) processPathProvider.Provider {
			return processPathProvider.NewProvider(pathProvider, pathGenerator)
		},
		func(pathProvider processPathProvider.Provider) processor.Factory {
			return processor.NewProcessorFactory(pathProvider)
		},
		func() process.Convertor { return process.NewConvertor() },
		func(urlProvider url.Provider) result.Factory { return result.NewFactory(urlProvider) },
		func(pathProvider path.PathProvider,
			urlProvider url.Provider,
			processesModel imageModel.ImageProcesses,
			resultFactory result.Factory,
		) runner.ProcessRunner {
			return runner.NewCachedRunner(
				runner.NewProcessRunner(pathProvider, urlProvider, processesModel, resultFactory),
				urlProvider,
				processesModel,
				resultFactory,
			)
		},
		func(model imageModel.Model, storage storage.Storage) Saver {
			return NewSaver(model, storage)
		},
	)
}
