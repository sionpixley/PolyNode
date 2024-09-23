import { DownloadedComponent } from './downloaded.component';

describe('DownloadedComponent', () => {
  let component: DownloadedComponent;

  beforeEach(() => {
    component = new DownloadedComponent();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should order downloadedVersions in desc', () => {
    component.downloadedVersions = ['v18.19.1', 'v13.15.1', 'v20.17.0'];
    component.currentVersion = 'v18.19.1';
    component.ngOnChanges();

    let expected = ['v20.17.0', 'v18.19.1', 'v13.15.1'];
    expect(component.downloadedVersions).toStrictEqual(expected);
  });
});
