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
	filenameGenerator := path.NewFilenameProvider()
	pathProvider := path.NewPathProvider(config.App.UploadPath)
	storage := storage.NewUploadStorage(filenameGenerator, pathProvider)
	processPathProvider := processPathProvider.NewProvider(pathProvider, filenameGenerator)

	di.InitModule("image",
		func() process.Convertor { return process.NewConvertor() },
		func() processor.Factory { return processor.NewProcessorFactory(processPathProvider) },
		func(urlProvider url.Provider) result.Factory { return result.NewFactory(urlProvider) },
		func(urlProvider url.Provider,
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
		func(model imageModel.Model) Saver { return NewSaver(model, storage) },
	)
}
