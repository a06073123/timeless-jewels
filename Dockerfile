FROM node:18-alpine AS base
# install pnpm
RUN  npm i -g pnpm
# install dependencies package
FROM base AS dependencies
WORKDIR /dependencies
COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN  pnpm install
# build
FROM base AS build
WORKDIR /build
COPY frontend/ .
COPY --from=dependencies /dependencies/node_modules ./node_modules
# deploy
FROM base AS deploy
WORKDIR /app
COPY --from=build /build/ /app/
ENTRYPOINT [ "pnpm","run","preview","--host" ]