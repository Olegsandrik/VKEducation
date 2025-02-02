import os
import time

import pytest
from _pytest.fixtures import FixtureRequest
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait

from Selenium.code.ui.locators import vk_locators
from Selenium.code.ui.pages.base_page_VK import BasePageVK
from ui.pages.base_page import BasePage
from Selenium.code.ui.url.urls import ConfigUrls
from selenium.webdriver.support import expected_conditions as EC
from dotenv import load_dotenv
load_dotenv()


class BaseCase:
    authorize = True

    @pytest.fixture(scope='function', autouse=True)
    def setup(self, driver, config, request: FixtureRequest):
        self.driver = driver
        self.config = config

        self.login_page = LoginPage(driver)


@pytest.fixture(scope='session')
def credentials():
    return {
        'user': os.getenv('EMAIL'),
        'password': os.getenv('PASSWORD')
    }

class LoginPage(BasePageVK):
    url = ConfigUrls.BASE

    def login(self, credentials):
        self.driver.maximize_window()
        self.click(
            vk_locators.LoginPageLocators.GO_BUTTON_AUTHBUTTON_LOCATOR, timeout=10
        )
        self.click(
            vk_locators.LoginPageLocators.GO_BUTTON_AUTH_WHITHOUT_VK_LOCATOR, timeout=10
        )
        self.input(
            vk_locators.LoginPageLocators.LOGIN_INPUT_LOCATOR,
            credentials.get('user', ''),
        )
        self.input(
            vk_locators.LoginPageLocators.PASSWORD_INPUT_LOCATOR,
            credentials.get('password', '')
        )
        self.click(
            vk_locators.LoginPageLocators.GO_BUTTON_SUBMIT_LOCATOR, timeout=10
        )


class MainPage(BasePageVK):
    url = ConfigUrls.FEED
    def open_main(self):
        self.driver.get(ConfigUrls.FEED)
        WebDriverWait(self.driver, 10).until(
            EC.visibility_of_element_located((By.TAG_NAME, 'body'))
        )

    def open_schedule(self):
        self.driver.get(ConfigUrls.SCHEDULE)
        WebDriverWait(self.driver, 10).until(
            EC.visibility_of_element_located((By.TAG_NAME, 'body'))
        )
        # time.sleep(1) # без этого кнопка есть, прожимается, но при этом скрипт не работает.
        # Кнопка жмется, а далее ничего не происходит

    def open_url(self, url):
        self.driver.get(url)
        WebDriverWait(self.driver, 10).until(
            EC.visibility_of_element_located((By.TAG_NAME, 'body'))
        )


class TestLogin(BaseCase):
    authorize = True

    def test_login(self, credentials):
        pass


class TestLK(BaseCase):

    def test_login(self, credentials):
        login_page = LoginPage(self.driver)
        login_page.login(credentials)

    def test_friend(self, credentials):
        login_page = LoginPage(self.driver)
        login_page.login(credentials)

        main_page = MainPage(self.driver)

        profile = main_page.find(
            vk_locators.MainPageLocators.PROFILE_LOCATOR, timeout=10
        )
        href_profile = profile.get_attribute('href')
        main_page.open_url(href_profile)

        friends = main_page.find(
            vk_locators.MainPageLocators.FRIENDS_LOCATOR, timeout=10
        )
        href_friends = friends.get_attribute('href')

        main_page.open_url(href_friends)

        first_friend = main_page.find(
            vk_locators.MainPageLocators.ALL_FRIENDS_LOCATOR
        )
        username = first_friend.text  # Искомый юзернейм
        print("\nusername: " + username)

    def test_lesson(self, credentials):
        login_page = LoginPage(self.driver)
        login_page.login(credentials)

        main_page = MainPage(self.driver)
        main_page.open_schedule()

        WebDriverWait(self.driver, 10).until(
            EC.presence_of_element_located(
                vk_locators.MainPageLocators.LOGOUT_SCRIPT
            )
        )

        main_page.click(
            vk_locators.MainPageLocators.GO_BOTTON_SEMESTR_LOCATOR, timeout=10
        )

        seminar = main_page.find(
            vk_locators.MainPageLocators.SEMINAR_INFO_LOCATOR
        )
        href_seminar = seminar.get_attribute('href')

        main_page.open_url(href_seminar)

        description = main_page.find_all(
            vk_locators.MainPageLocators.DESCRIPTION_LOCATOR
        )[1]
        print("\nlesson description: " + description.text)

