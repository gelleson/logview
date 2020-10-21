import { Injectable } from '@angular/core';
import { ReaderService } from '../../api/reader.api';
import { from, Observable } from 'rxjs';
import { Line } from '../../models/line';
import { LogInfo } from '../../models/logInfo';

declare var window: any;

@Injectable({
    providedIn: 'root',
})
export class ReaderServiceImpl implements ReaderService {
    constructor() {}

    match(logName: string, query: string): Observable<Line[]> {
        return from<Array<Line[]>>(
            window.backend.LogService.Match(logName, query)
        );
    }

    offset(logName: string, limit: number, offset: number): Observable<Line[]> {
        return from<Array<Line[]>>(
            window.backend.LogService.Offset(logName, limit, offset)
        );
    }

    logList(): Observable<LogInfo[]> {
        return from<Array<LogInfo[]>>(window.backend.LogService.LogsList());
    }
}
