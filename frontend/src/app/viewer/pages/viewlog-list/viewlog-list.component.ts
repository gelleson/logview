import { Component, OnInit } from '@angular/core';
import { ReaderService } from '../../../core/services/api/reader.api';
import { UploaderService } from '../../../core/services/api/uploader.api';
import { LogInfo } from '../../../core/services/models/logInfo';

@Component({
    selector: 'app-viewlog-list',
    templateUrl: './viewlog-list.component.html',
    styleUrls: ['./viewlog-list.component.css'],
})
export class ViewlogListComponent implements OnInit {
    logs: LogInfo[] = [];
    constructor(
        private readerService: ReaderService,
        private uploadService: UploaderService
    ) {}

    ngOnInit() {
        this.readerService.logList().subscribe((logs) => {
            if (!logs.length) {
                // this.uploadService
                //   .subscribe();
            }
            this.logs = logs;
        });
    }
}
