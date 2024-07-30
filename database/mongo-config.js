const mongoose = require('mongoose');

mongoose.connect('mongodb://localhost:27017/flightstatus', {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

const flightSchema = new mongoose.Schema({
  id: Number,
  status: String,
});

const Flight = mongoose.model('Flight', flightSchema);

module.exports = Flight;
