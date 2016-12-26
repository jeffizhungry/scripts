#!/usr/bin/env python3.4

import time

from selenium import webdriver
from selenium.webdriver.common.keys import Keys

SITE = "https://coachella-weekend2.frontgatetickets.com/coachella2/captcha.html"
SITE = "https://www.coachella.com"

if __name__ == "__main__":
    driver = webdriver.Firefox()
    driver.get(SITE)

    while True:
        print("looping...")

         #         # Note: When selenium can't find the link, an exception is thrown.
         #         #       For our purposes, this means clicking the link was successful
         #         #       in the last iteration, and we have nagivated away. So, the
         #         #       script will exit here and the brower will be left open for
         #         #       a human user to enter the captcha.
         #         link = driver.find_element_by_link_text('click here')
         #
         #         link.click()
         #         time.sleep(5)

        driver.refresh()
        time.sleep(10)
