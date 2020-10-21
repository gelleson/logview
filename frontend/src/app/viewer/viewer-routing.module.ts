import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ViewerComponent } from './viewer.component';
import { ViewlogListComponent } from './pages/viewlog-list/viewlog-list.component';
import { ViewlogDetailComponent } from './pages/viewlog-detail/viewlog-detail.component';

const routes: Routes = [
    {
        path: '',
        component: ViewerComponent,
        children: [
            {
                path: '',
                component: ViewlogListComponent,
            },
            {
                path: ':id',
                component: ViewlogDetailComponent,
            },
        ],
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class ViewerRoutingModule {}
