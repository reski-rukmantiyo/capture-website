# README.md

## How to Run

1. Clone the repository
2. Install Docker
3. Run the following command in the terminal:

```bash
docker run -d -p 9222:9222 --rm --name headless-shell --shm-size 2G chromedp/headless-shell
```
4. Create file config.env with the conent

```
DOCKER_URL=wss://localhost:9222
```

6. Run folllowing command

```bash
./scrapper https://www.google.com
```
