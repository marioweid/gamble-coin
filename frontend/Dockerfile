# Use the official Node.js image to build the frontend
FROM node:22.9 AS build

WORKDIR /app
COPY . .

# Clean install all node modules
RUN npm ci
RUN npm run build

# Clean Container
RUN rm -rf src/ static/ emailTemplates/ docker-compose.yml

# Run Command
USER node:node
CMD ["node","--env-file=.env", "build/index.js"]