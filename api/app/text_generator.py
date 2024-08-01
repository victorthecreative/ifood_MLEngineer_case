from transformers import GPT2LMHeadModel, GPT2Tokenizer
from pathlib import Path
import sys


def load_model(model_path):
    model = GPT2LMHeadModel.from_pretrained(model_path, local_files_only=True)
    return model


def load_tokenizer(tokenizer_path):
    tokenizer = GPT2Tokenizer.from_pretrained(
        tokenizer_path, local_files_only=True
    )
    tokenizer.pad_token = tokenizer.eos_token
    return tokenizer


def generate_text(sequence, max_length, model, tokenizer):
    ids = tokenizer.encode(f"{sequence}", return_tensors="pt")
    final_outputs = model.generate(
        ids,
        do_sample=True,
        max_length=max_length,
        pad_token_id=model.config.eos_token_id,
        top_k=50,
        top_p=0.95,
    )
    return tokenizer.decode(final_outputs[0], skip_special_tokens=True)


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python3 text_generator.py <sequence> <max_length>")
        sys.exit(1)

    sequence = sys.argv[1]
    max_length = int(sys.argv[2])
    model_path = Path("/app/api/model_trained")

    model = load_model(model_path)
    tokenizer = load_tokenizer(model_path)

    result = generate_text(sequence, max_length, model, tokenizer)
    print(result)
