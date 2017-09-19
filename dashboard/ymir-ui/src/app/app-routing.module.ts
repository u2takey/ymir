import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { JobListComponent } from './job/job.list';
import { JobDetailComponent } from './job/job.detail';
import { JobEditComponent } from './job/job.edit';


const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard', component: JobListComponent }, // joblist
  { path: 'jobview/:jobid', component: JobDetailComponent },  // jobdetail, resultlist
  { path: 'jobedit/:jobid', component: JobEditComponent },   // jobnew
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
