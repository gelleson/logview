import {Observable} from 'rxjs';
import {Injectable} from '@angular/core';
import {Line} from '../models/line';
import {LogInfo} from '../models/logInfo';

@Injectable()
export abstract class ReaderService {
  public abstract offset(logName: string, limit: number, offset: number): Observable<Line[]>;
  public abstract match(logName: string, query: string): Observable<Line[]>;
  public abstract logList(): Observable<LogInfo[]>;
}
