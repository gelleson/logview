import { ReaderService } from './reader.api';
import { UploaderService } from './uploader.api';

export interface Backend {
    LogService: ReaderService;
    UploadService: UploaderService;
}
