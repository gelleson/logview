import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewlogListComponent } from './viewlog-list.component';

describe('ViewlogListComponent', () => {
    let component: ViewlogListComponent;
    let fixture: ComponentFixture<ViewlogListComponent>;

    beforeEach(async(() => {
        TestBed.configureTestingModule({
            declarations: [ViewlogListComponent],
        }).compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(ViewlogListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
