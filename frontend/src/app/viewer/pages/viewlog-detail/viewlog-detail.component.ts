import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ReaderService } from '../../../core/services/api/reader.api';
import { Line } from '../../../core/services/models/line';
import { FormBuilder, FormGroup } from '@angular/forms';
import { debounceTime } from 'rxjs/operators';

@Component({
    selector: 'app-viewlog-detail',
    templateUrl: './viewlog-detail.component.html',
    styleUrls: ['./viewlog-detail.component.css'],
})
export class ViewlogDetailComponent implements OnInit {
    id = this.routerSnapshot.snapshot.paramMap.get('id');
    lines: Line[] = [];
    form: FormGroup;
    constructor(
        private routerSnapshot: ActivatedRoute,
        private readerService: ReaderService,
        fb: FormBuilder
    ) {
        this.form = fb.group({
            search: null,
        });

        this.form
            .get('search')
            .valueChanges.pipe(debounceTime(100))
            .subscribe((term) => {
                if (!term || term === '') {
                    this.readerService
                        .offset(this.id, 100, 0)
                        .subscribe((resp) => {
                            this.lines = resp;
                        });
                } else {
                    this.readerService
                        .match(this.id, term)
                        .subscribe((logs) => (this.lines = logs));
                }
            });
    }

    ngOnInit() {
        this.readerService.offset(this.id, 100, 0).subscribe((resp) => {
            this.lines = resp;
        });
    }
}
