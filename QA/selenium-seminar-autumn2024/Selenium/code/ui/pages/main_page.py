from selenium.webdriver.common.by import By

import allure
from Selenium.code.ui.locators import basic_locators
from Selenium.code.ui.pages.base_page import BasePage
from Selenium.code.ui.pages.events_page import EventsPage


class MainPage(BasePage):

    locators = basic_locators.MainPageLocators()

    @allure.step("Step 2")
    def go_to_events_page(self):
        events_button = self.find(self.locators.EVENTS)
        # self.click(events_button)
        self.click((By.ID, 'events'))
        return EventsPage(self.driver)
