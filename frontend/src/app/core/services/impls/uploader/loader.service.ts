import {Injectable} from '@angular/core';
import {from, Observable} from 'rxjs';
import {UploaderService} from '../../api/uploader.api';
import {Line} from '../../models/line';

declare var window: any;

@Injectable({
  providedIn: 'root'
})
export class LoaderServiceImpl implements UploaderService {
  public static demoPath = '/Users/user/go/src/github.com/gelleson/logview/000000';

  constructor() {
  }

  public upload(filename: string): Observable<any> {
    return from<any>(
      window.backend.UploadService.UploadFile(filename),
    );
  }
}
