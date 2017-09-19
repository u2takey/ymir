import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { AppComponent } from './app.component';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { JobService } from './service/job';
import { JobListComponent } from './job/job.list';
import { JobEditComponent } from './job/job.edit';
import { JobDetailComponent } from './job/job.detail';
import { NgxDatatableModule } from '@swimlane/ngx-datatable';
import { AppRoutingModule } from './app-routing.module';
import { MdButtonModule, MdToolbarModule, MdSidenavModule, MdListModule } from '@angular/material';
import { MdInputModule, MdCheckboxModule } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MdCardModule } from '@angular/material';

@NgModule({
  declarations: [
    AppComponent,
    JobListComponent,
    JobEditComponent,
    JobDetailComponent,
  ],
  imports: [
    NgbModule.forRoot(),
    BrowserModule,
    AppRoutingModule,
    HttpModule,
    FormsModule,
    MdButtonModule,
    MdToolbarModule,
    MdSidenavModule,
    MdInputModule,
    MdCheckboxModule,
    MdListModule,
    MdCardModule,
    NgxDatatableModule,
    BrowserAnimationsModule
  ],
  providers: [JobService],
  bootstrap: [AppComponent]
})
export class AppModule { }
