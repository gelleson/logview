import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewlogDetailComponent } from './viewlog-detail.component';

describe('ViewlogDetailComponent', () => {
    let component: ViewlogDetailComponent;
    let fixture: ComponentFixture<ViewlogDetailComponent>;

    beforeEach(async(() => {
        TestBed.configureTestingModule({
            declarations: [ViewlogDetailComponent],
        }).compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(ViewlogDetailComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
