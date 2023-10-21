FROM node:18

WORKDIR /app

COPY package*.json ./

RUN npm install
RUN apt update
RUN apt install -y libgtk2.0-0 libgtk-3-0 libgbm-dev libnotify-dev libgconf-2-4 libnss3 libxss1 libasound2 libxtst6 xauth xvfb

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]