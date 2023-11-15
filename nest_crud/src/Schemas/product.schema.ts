import mongoose from 'mongoose';

export const ProductSchema = new mongoose.Schema({
  name: String,
  description: String,
  qty: Number,
  price: Number,
});

ProductSchema.set('toJSON', {
  transform: (doc, returnedDoc) => {
    (returnedDoc.id = returnedDoc._id.toString()),
      delete returnedDoc._id,
      delete returnedDoc.__v;
  },
});
