# Node base image
FROM node:12.3.0

# set lable
LABEL key="Hammad Tahir <hammad.tahir.ch@gmail.com>"

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY ./client/package.json /app/package.json

# install npm packages
RUN npm install --silent

# install react scripts
RUN npm install react-scripts@3.0.1 -g --silent

# expose port 2001 for public access
EXPOSE 2001

# start app
CMD ["npm", "run","dev"]