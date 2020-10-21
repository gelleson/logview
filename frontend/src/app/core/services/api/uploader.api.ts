import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';

@Injectable()
export abstract class UploaderService {
    public abstract upload<T>(filename: string): Observable<T>;
}
