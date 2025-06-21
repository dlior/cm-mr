console.log('App 1: Main script executed');
fetch('https://api.chucknorris.io/jokes/random')
  .then((response) => response.json())
  .then(({ value }) => console.log('joke:', value))
  .catch(console.error);
