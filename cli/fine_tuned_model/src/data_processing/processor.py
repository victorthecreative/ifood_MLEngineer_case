import pandas as pd
from data_cleaning.cleaner import DataCleaner


class ArticleProcessor:
    def __init__(self, csv_path, txt_path, encoding="ISO-8859-1"):
        self.csv_path = csv_path
        self.txt_path = txt_path
        self.encoding = encoding

    def process_articles(self):
        df = pd.read_csv(self.csv_path, encoding=self.encoding)
        df = df.dropna()
        with open(self.txt_path, "w") as text_data:
            for _, item in df.iterrows():
                cleaner = DataCleaner(item["Article"])
                article = cleaner.clean()
                text_data.write(article)
