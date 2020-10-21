import { ModuleWithProviders, NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReaderServiceImpl } from './services/impls/reader/reader.service';
import { ReaderService } from './services/api/reader.api';
import { LoaderServiceImpl } from './services/impls/uploader/loader.service';
import { UploaderService } from './services/api/uploader.api';

@NgModule({
    declarations: [],
    imports: [CommonModule],
})
export class CoreModule {
    public static forRoot(): ModuleWithProviders {
        return {
            ngModule: CoreModule,
            providers: [
                {
                    provide: ReaderService,
                    useClass: ReaderServiceImpl,
                },
                {
                    provide: UploaderService,
                    useClass: LoaderServiceImpl,
                },
            ],
        };
    }
}
