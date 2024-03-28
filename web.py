import requests
from bs4 import BeautifulSoup
import random

class webScraper:   
        def webScrape():
                global quote
                url = "https://www.goodreads.com/quotes/tag/stoicism?page=2"
                page = requests.get(url)
                soup = BeautifulSoup(page.content, "html.parser")
                quotes = soup.find_all("div", class_="quoteText")
                quoteS = random.choice(quotes)
                quote = quoteS.text
                quote = quote.replace("“", "")
                quote = quote.replace("”", "")
                quote = quote.replace("", "") 
                
        webScrape()

class test:
        def duplicated(): # if the quote is duplicated the function will loop again until it finds a new quote
                while quote != quote:
                        print("Quote is duplicated")
                        return
                else:
                        quote == quote # if the quote is not duplicated the function will stop
                        webScraper.webScrape()
                        print(quote)
        duplicated()
        def tooLong(): # if the quote is too long the function will loop again until it finds a new quote
                while len(quote) > 280:
                        return
                else:
                        len(quote) < 280
                        webScraper.webScrape
        tooLong()


# main function that will run the program and break
def main():
    main()