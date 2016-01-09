import time

from selenium import webdriver
from selenium.webdriver.common.keys import Keys

SITE = "https://coachella-weekend2.frontgatetickets.com/coachella2/captcha.html"

if __name__ == "__main__":
    driver = webdriver.Firefox()
    driver.get(SITE)

    while True:
        print("looping...")

        # Note: When selenium can't find the link, an exception is thrown.
        #       For our purposes, this means clicking the link was successful
        #       in the last iteration, and the script will exit right here.
        #       So, the browser will be kept open and the user can enter the
        #       captcha.
        link = driver.find_element_by_link_text('click here')

        link.click()
        time.sleep(5)

        driver.refresh()
        time.sleep(.5)
