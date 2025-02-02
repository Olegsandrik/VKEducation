from selenium.webdriver.common.by import By


class LoginPageLocators:
    GO_BUTTON_AUTHBUTTON_LOCATOR = (By.XPATH, "//a[text()='вход / регистрация']")
    GO_BUTTON_AUTH_WHITHOUT_VK_LOCATOR = (By.XPATH, "//button[text()='Продолжить с помощью почты и пароля']")
    LOGIN_INPUT_LOCATOR = (By.ID, 'email')
    PASSWORD_INPUT_LOCATOR = (By.ID, 'password')
    GO_BUTTON_SUBMIT_LOCATOR = (By.XPATH, "//button[text()='Войти с паролем']")

class MainPageLocators:
    PROFILE_LOCATOR = (By.XPATH, "//a[@class='full_name']")
    FRIENDS_LOCATOR = (By.XPATH, "//a[text()='Друзья']")
    ALL_FRIENDS_LOCATOR = (By.XPATH, "(//div[contains(@class, 'friends_item')])//p[@class='username']/a")
    GO_BOTTON_SEMESTR_LOCATOR = (By.CSS_SELECTOR, 'li[intervalid="semester"]')
    GO_BOTTON_NEAR_LOCATOR = (By.CSS_SELECTOR, 'li[intervalid="near"]')
    LOGOUT_SCRIPT = (By.CLASS_NAME, 'logout')
    SEMINAR_INFO_LOCATOR = (By.CSS_SELECTOR, 'tr#schedule_item_1729544400 a.schedule-show-info')
    DESCRIPTION_LOCATOR = (By.XPATH, "//div[@class='description']//div[@class='section-text text']")