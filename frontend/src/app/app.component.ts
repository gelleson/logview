import { Component, OnInit } from '@angular/core';
import { UploaderService } from './core/services/api/uploader.api';
import { ReaderService } from './core/services/api/reader.api';
import { Line } from './core/services/models/line';
import { Form, FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { debounceTime } from 'rxjs/operators';

declare var window: any;

@Component({
    // tslint:disable-next-line:component-selector
    selector: '[id="app"]',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
    title = 'my-app';

    constructor() {}

    onClickMe() {}

    async ngOnInit() {}
}
