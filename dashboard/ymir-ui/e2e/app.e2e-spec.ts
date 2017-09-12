import { YmirUiPage } from './app.po';

describe('ymir-ui App', () => {
  let page: YmirUiPage;

  beforeEach(() => {
    page = new YmirUiPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!');
  });
});
