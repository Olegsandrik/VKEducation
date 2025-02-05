from selenium.webdriver.common.by import By

class LeadFormsLocators:
    CREATE_LEAD_FORM = (By.XPATH, "//button[contains(@class, 'LeadForms_createButton')]")
    CREATE_LANDING_SWITCH = (By.XPATH, "//div[contains(@class, 'LandingPageSelector_selectContent')]/label")
    NEW_LEADFORM_INPUTS = (By.XPATH, "//input[contains(@class, 'vkuiTypography vkuiInput__el')]")
    CONTINUE = (By.XPATH, "//button[@data-testid='submit']")
    ADD_LOGO = (By.XPATH, "//div[@data-testid='set-global-image']")
    LOAD_IMAGE = (By.XPATH, "//label[contains(@class, 'LocalFileSelector_file')]//input")
    LOAD_IMAGE_SIDEBAR = (By.XPATH, "//div[contains(@class, 'ModalSidebarPage_container')]")
    UPLOADED_IMAGE_ITEM = (By.XPATH, "//div[contains(@class, 'ImageItem_imageItem')]")
    ACTIVE_STEP = (By.XPATH, "//div[contains(@class, 'CreateLeadFormModal_activeStep')]/div[contains(@class, 'CreateLeadFormModal_stepCounter')]")
    FORM_TYPE = (By.XPATH, "//div[contains(@class, 'vkuiFormItem')]//label[contains(@class, 'vkuiSegmentedControlOption')]")
    BIG_FORM_DESCRIPTION = (By.XPATH, "//section[contains(@class, 'vkuiInternalGroup')]//textarea")
    DISCOUNT_TYPE = (By.XPATH, "//div[contains(@class, 'AwardBlock')]//div[@role='radiogroup']/label")
    DISCOUNT_INPUT = (By.XPATH, "//div[contains(@class, 'AwardBlock')]//input[@type='text']")
    ADD_BUTTON = (By.XPATH, "//section[contains(@class, 'vkuiInternalGroup')]//button[contains(@class, 'vkuiButton')]")
    QUESTION = (By.XPATH, "//div[contains(@class, 'Question_question__')]")
    REMOVE_QUESTION = (By.XPATH, "//button[@aria-label='Remove']")
    LEADFORM_ITEM = (By.XPATH, "//div[contains(@class, 'BaseTable__table-frozen-left')]//div[@class='BaseTable__row']")
    LEADFORM_ITEM_CHECKBOX = (By.XPATH, "//div[@role='gridcell']/div")
    LEADFORM_STATUS_DROPDOWN = (By.XPATH, "//div[contains(@class, 'LeadForms_selectStatus')]")