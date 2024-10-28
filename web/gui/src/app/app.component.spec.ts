import { TestBed } from "@angular/core/testing";
import { AppComponent } from "./app.component";
import { DownloadedComponent } from "./components/downloaded/downloaded.component";
import { AvailableComponent } from "./components/available/available.component";
import { SpinnerComponent } from "./components/spinner/spinner.component";
import { MatIconModule, MatIconRegistry } from "@angular/material/icon";
import { GuiService } from "./services/gui.service";
import { MatTableModule } from "@angular/material/table";
import { MatCheckboxModule } from "@angular/material/checkbox";
import { MatButtonModule } from "@angular/material/button";
import { MatTooltipModule } from "@angular/material/tooltip";
import { MatInputModule } from "@angular/material/input";
import { MatFormFieldModule } from "@angular/material/form-field";
import { ReactiveFormsModule } from "@angular/forms";
import { provideHttpClientTesting } from "@angular/common/http/testing";
import { provideHttpClient } from "@angular/common/http";
import { of } from "rxjs";

describe('AppComponent', () => {
  let component: AppComponent;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        AppComponent,
        DownloadedComponent,
        AvailableComponent,
        SpinnerComponent,
        MatTableModule,
        MatCheckboxModule,
        MatIconModule,
        MatButtonModule,
        MatTooltipModule,
        MatInputModule,
        MatFormFieldModule,
        ReactiveFormsModule
      ],
      providers: [MatIconRegistry, GuiService, provideHttpClient(), provideHttpClientTesting()]
    }).compileComponents();

    component = TestBed.createComponent(AppComponent).componentInstance;
  });

  afterEach(() => jest.clearAllMocks());

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should select LTS versions', () => {
    component['_api'].version = jest.fn(() => of('v0.0.0'));
    component['_api'].search = jest.fn(() => of(['v10.13.1', '11.33.1 (lts)', '12.1.22 (lts)', '13.4.4']));

    let expected = ['11.33.1', '12.1.22'];
    component.ngOnInit();
    expect(component.ltsVersions).toStrictEqual(expected);
  });
});
