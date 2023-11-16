import mongoose from 'mongoose';

export const ProductSchema = new mongoose.Schema({
  name: {
    type: String,
    required: true,
  },
  description: String,
  qty: {
    type: Number,
    default: 0,
  },
  price: {
    type: Number,
    required: true,
  },
});

ProductSchema.set('toJSON', {
  transform: (doc, returnedDoc) => {
    (returnedDoc.id = returnedDoc._id.toString()),
      delete returnedDoc._id,
      delete returnedDoc.__v;
  },
});
