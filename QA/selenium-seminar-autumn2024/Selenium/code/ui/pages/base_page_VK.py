import time

import allure
from selenium.common import TimeoutException
from selenium.webdriver.remote.webelement import WebElement
from Selenium.code.ui.locators import basic_locators, vk_locators
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from Selenium.code.ui.url.urls import ConfigUrls


class PageNotOpenedExeption(Exception):
    pass


class BasePageVK(object):

    locators = vk_locators.LoginPageLocators()
    url = ConfigUrls.BASE

    def is_opened(self, timeout=15):
        started = time.time()
        while time.time() - started < timeout:
            if self.driver.current_url == self.url:
                return True
        raise PageNotOpenedExeption(f'{self.url} did not open in {timeout} sec, current url {self.driver.current_url}')

    def __init__(self, driver):
        self.driver = driver
        self.is_opened()

    def wait(self, timeout=None):
        if timeout is None:
            timeout = 5
        return WebDriverWait(self.driver, timeout=timeout)

    @allure.step('find one element')
    def find(self, locator, timeout=None):
        try:
            return self.wait(timeout).until(EC.presence_of_element_located(locator))
        except TimeoutException:
            print(f'TimeoutException. Item was not found')
            return None

    @allure.step('find all elements')
    def find_all(self, locator, timeout=None):
        return self.wait(timeout).until(EC.presence_of_all_elements_located(locator))

    @allure.step('input data in element')
    def input(self, input_field, data):
        elem = self.find(input_field)
        elem.clear()
        elem.send_keys(data)

    @allure.step('click on element')
    def click(self, locator, timeout=None) -> WebElement:
        elem = self.find(locator, timeout=timeout)
        elem.click()
