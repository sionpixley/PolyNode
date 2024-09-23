import { AvailableComponent } from "./available.component";

describe('AvailableComponent', () => {
  let component: AvailableComponent;

  beforeEach(() => {
    component = new AvailableComponent();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should order availableVersions in desc', () => {
    component.availableVersions = ['v18.11.3', 'v18.2.1', 'v10.12.1', 'v20.22.1', 'v8.19.1'];
    component.ngOnChanges();

    let expected = ['v20.22.1', 'v18.11.3', 'v18.2.1', 'v10.12.1', 'v8.19.1'];
    expect(component.availableVersions).toStrictEqual(expected);
  });

  it('should select lts versions correctly', () => {
    component.availableVersions = [
      'v22.9.0',
      'v21.7.3',
      'v19.9.0',
      'v17.9.1',
      'v15.14.0',
      'v13.14.0',
      'v11.15.0',
      'v20.17.0',
      'v18.20.4',
      'v16.20.2',
      'v14.21.3',
      'v12.22.12',
      'v10.24.1',
      'v8.17.0'
    ];
    component.ngOnChanges();

    let expected = [
      'v20.17.0',
      'v18.20.4',
      'v16.20.2',
      'v14.21.3',
      'v12.22.12',
      'v10.24.1',
      'v8.17.0'
    ];
    expect(component['_ltsVersions']).toStrictEqual(expected);
  });
});
