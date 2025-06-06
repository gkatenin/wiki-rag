## 🔍 Key Features

This project is a minimal, efficient Retrieval-Augmented Generation (RAG) system written entirely in Go. It emphasizes simplicity, speed, and portability, avoiding heavy dependencies or external frameworks.

### 1. 🌐 Lightweight and Efficient Wikipedia RAG

- Fetches Wikipedia articles on the fly using the standard Go `net/http` package.
- No external APIs or scraping tools required — just clean and direct HTTP requests.
- Multilingual support enables access to Wikipedia content in different languages (e.g., `en`, `de`, etc.).

### 2. 🔗 Native Support for BERT Embedding Models

- Utilizes the lightweight [`github.com/kelindar/search`](https://github.com/kelindar/search) package for semantic search and embedding.
- Embeddings are powered by [llama.cpp](https://github.com/ggerganov/llama.cpp), integrated seamlessly using the [purego](https://github.com/ebitengine/purego) library.
- This design eliminates the need for `cgo`, making the project fully self-contained and easier to build and deploy across platforms.

### 3. 📦 Chunking and Semantic Indexing

- Articles are split into overlapping chunks to preserve contextual flow (default: 1600 characters with 100-character overlap).
- Each chunk is embedded and indexed using vector search, enabling fast semantic retrieval of the most relevant content.
- The result is a lightweight, high-performance context-building system ideal for downstream LLM use.

### 4. 🧠 GGUF Model Compatibility

- Compatible with a wide variety of sentence-level BERT-style models in the [GGUF format](https://github.com/ggerganov/ggml/blob/master/docs/gguf.md), such as `paraphrase-multilingual-MiniLM-*`.
- You can easily swap or experiment with different language models as long as they follow the GGUF standard.

## 📚 How to Use

```sh
% ./wiki-rag -query "otto klemperer in the usa" -lang en
....................
Otto Klemperer in the USA refers to the period when Klemperer, a German-born conductor and composer, lived and worked in the United States. Klemperer was widely regarded as one of the leading conductors of the 20th century. He fled Germany in 1933 due to the rise of the Nazi Party and settled in the United States, where he had a significant impact on American classical music.

Here are some key points about Otto Klemperer's time in the USA:

1. **Career in the USA**: After leaving Germany, Klemperer first settled in the United States in 1933. He initially worked as a guest conductor for various orchestras before being appointed as the music director of the Los Angeles Philharmonic from 1933 to 1939.
2. **Los Angeles Philharmonic**: During his tenure with the Los Angeles Philharmonic, Klemperer was instrumental in shaping the orchestra's musical direction and reputation. He brought a high level of musical sophistication and artistic excellence to the orchestra.
3. **Pittsburgh Symphony Orchestra**: Following his time in Los Angeles, Klemperer returned to the United States to take up the position of music director of the Pittsburgh Symphony Orchestra from 1946 to 1952. During this period, he worked with the orchestra to improve its musical standards and expanded its repertoire.
4. **Later Years**: After leaving the Pittsburgh Symphony Orchestra, Klemperer continued to conduct and was in great demand as a guest conductor for various orchestras around the world. He also worked as a conductor and music director for the San Francisco Symphony from 1952 to 1954.
Klemperer's time in the USA was marked by his dedication to music education and his efforts to promote musical excellence, making him a pivotal figure in the development of American classical music during the mid-20th century.

% ./wiki-rag -query "сергей рахманинов в щвейцарии" -lang ru
....................
1. Сергей Рахманинов, выдающийся русский композитор и пианист, провел значительную часть своей жизни за границей.
2. Он эмигрировал из России в 1917 году, после Октябрьской революции, и жил в основном в США, но также часто бывал в Европе.
3. В 1931 году Рахманинов переехал в Швейцарию, где провел последние годы своей жизни.
4. Швейцария стала для него местом покоя и вдохновения, где он продолжал сочинять музыку и выступать на концертах.

Исходя из данного контекста, "Сергей Рахманинов в Швейцарии" можно охарактеризовать следующим образом:

Сергей Рахманинов провел последние годы своей жизни в Швейцарии, где нашел спокойствие и вдохновение для творчества. Эмигрировав из России в 1917 году, он обосновался в США, но в 1931 году переехал в Швейцарию, которая стала его последним домом. В Швейцарии Рахманинов продолжал сочинять музыку и давать концерты, внося значительный вклад в мировую музыкальную культуру. Этот период его жизни был отмечен как профессиональным ростом, так и личным благополучием, что позволило ему завершить некоторые из своих наиболее известных произведений, включая «Рапсодию на тему Паганини».
```
