from Selenium.code.ui.locators import basic_locators
from Selenium.code.ui.pages.base_page import BasePage


class EventsPage(BasePage):

    locators = basic_locators.EventsPageLocators()
    url = 'https://www.python.org/events/'
