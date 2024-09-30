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
});
