from data_processing.processor import ArticleProcessor
from model_training.trainer import GPT2Trainer
import yaml
from pathlib import Path
import shutil
import os


def load_config(config_file):
    with open(config_file, "r") as file:
        config = yaml.safe_load(file)
    return config


if __name__ == "__main__":

    output_path = (
        Path(__file__).parent.parent.parent.parent / "api" / "model_trained"
    )

    for item in output_path.iterdir():
        if item.is_dir():
            shutil.rmtree(item)

    for item in output_path.iterdir():
        if item.is_file():
            item.unlink()

    config_path = Path(__file__).parent / "config.yml"

    if not config_path.exists():
        raise FileNotFoundError(f"Config file not found: {config_path}")

    config = load_config(config_path)

    os.chdir(Path(__file__).parent.parent)

    raw_data_path = Path(config["train_file_path_raw"]).resolve()
    curated_data_path = Path(config["train_file_path_curated"]).resolve()

    if not raw_data_path.exists():
        raise FileNotFoundError(f"Raw data file not found: {raw_data_path}")

    processor = ArticleProcessor(
        csv_path=raw_data_path,
        txt_path=curated_data_path,
    )

    processor.process_articles()

    trainer = GPT2Trainer(
        model_name=config["model_id"],
        output_dir=output_path,
        overwrite_output_dir=config["overwrite_output_dir"],
        per_device_train_batch_size=config["per_device_train_batch_size"],
        num_train_epochs=config["num_train_epochs"],
        save_steps=config["save_steps"],
    )
    trainer.train(Path(config["train_file_path_curated"]))
