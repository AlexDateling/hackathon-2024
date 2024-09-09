# Build stage
FROM node:14-buster-slim AS build

WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json* ./
RUN npm ci

# Copy the rest of the application and build
COPY . .
RUN npm run build

# Runtime stage
FROM python:3.9-slim-buster

# Install Node.js
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && curl -sL https://deb.nodesource.com/setup_14.x | bash - \
    && apt-get install -y --no-install-recommends \
    nodejs \
    && pip install --no-cache-dir --upgrade pip pipenv \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /api

# Copy built assets from the build stage
COPY --from=build /app/dist ./dist
COPY --from=build /app/package.json ./

# Install only production dependencies
RUN npm ci --only=production

# Copy any additional necessary files (e.g., Python files)
COPY *.py ./

EXPOSE 3000

CMD ["npm", "run", "start:prod"]
