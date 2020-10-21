import { TestBed } from '@angular/core/testing';

import { LoaderService } from './uploader.service';

describe('UploaderService', () => {
    beforeEach(() => TestBed.configureTestingModule({}));

    it('should be created', () => {
        const service: LoaderService = TestBed.get(LoaderService);
        expect(service).toBeTruthy();
    });
});
