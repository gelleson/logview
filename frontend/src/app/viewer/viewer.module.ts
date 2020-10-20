import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ViewerRoutingModule } from './viewer-routing.module';
import { ViewerComponent } from './viewer.component';
import { ViewlogDetailComponent } from './pages/viewlog-detail/viewlog-detail.component';
import { ViewlogListComponent } from './pages/viewlog-list/viewlog-list.component';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";

@NgModule({
  declarations: [ViewerComponent, ViewlogDetailComponent, ViewlogListComponent,],
  imports: [
    CommonModule,
    ViewerRoutingModule,
    FormsModule,
    ReactiveFormsModule,
  ]
})
export class ViewerModule { }
