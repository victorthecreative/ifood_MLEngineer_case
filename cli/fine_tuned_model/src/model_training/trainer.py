from transformers import TextDataset, DataCollatorForLanguageModeling
from transformers import GPT2Tokenizer, GPT2LMHeadModel
from transformers import Trainer, TrainingArguments


class GPT2Trainer:
    def __init__(
        self,
        model_name,
        output_dir,
        overwrite_output_dir,
        per_device_train_batch_size,
        num_train_epochs,
        save_steps,
    ):
        self.model_name = model_name
        self.output_dir = output_dir
        self.overwrite_output_dir = overwrite_output_dir
        self.per_device_train_batch_size = per_device_train_batch_size
        self.num_train_epochs = num_train_epochs
        self.save_steps = save_steps
        self.tokenizer = GPT2Tokenizer.from_pretrained(model_name)
        self.model = GPT2LMHeadModel.from_pretrained(model_name)

    def load_dataset(self, file_path, block_size=128):
        dataset = TextDataset(
            tokenizer=self.tokenizer,
            file_path=file_path,
            block_size=block_size,
        )
        return dataset

    def load_data_collator(self, mlm=False):
        data_collator = DataCollatorForLanguageModeling(
            tokenizer=self.tokenizer,
            mlm=mlm,
        )
        return data_collator

    def train(self, train_file_path):
        train_dataset = self.load_dataset(train_file_path)
        data_collator = self.load_data_collator()

        training_args = TrainingArguments(
            output_dir=self.output_dir,
            overwrite_output_dir=self.overwrite_output_dir,
            per_device_train_batch_size=self.per_device_train_batch_size,
            num_train_epochs=self.num_train_epochs,
            save_steps=self.save_steps,
        )

        trainer = Trainer(
            model=self.model,
            args=training_args,
            data_collator=data_collator,
            train_dataset=train_dataset,
        )

        trainer.train()
        trainer.save_model()

        self.tokenizer.save_pretrained(self.output_dir)
        self.model.save_pretrained(self.output_dir)
