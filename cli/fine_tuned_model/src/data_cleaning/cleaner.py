import re


class DataCleaner:
    def __init__(self, text):
        self.text = str(text)

    def clean(self):
        self.text = re.sub(r"\s\W", " ", self.text)
        self.text = re.sub(r"\W,\s", " ", self.text)
        self.text = re.sub(r"\d+", "", self.text)
        self.text = re.sub(r"\s+", " ", self.text)
        self.text = re.sub(r"[!@#$_]", "", self.text)
        self.text = self.text.replace("co", "")
        self.text = self.text.replace("https", "")
        self.text = self.text.replace(r"[\w*", " ")
        return self.text
