import vertexai
from vertexai.generative_models import GenerativeModel

# Initialize
vertexai.init(project="bakery-street-project", location="us-central1")

# Use Gemini 1.5 Flash
model = GenerativeModel("gemini-1.5-flash")

# Test call
try:
    response = model.generate_content("Hello! Are you active on the bakery-street-project?")
    print(f"\nSuccess! Response:\n{response.text}")
except Exception as e:
    print(f"\nError: {e}")
