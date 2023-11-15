import { Injectable } from '@nestjs/common';
import { Product } from './interface/products.interface';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
@Injectable()
export class ProductService {
  constructor(
    @InjectModel('products') private readonly productModel: Model<Product>,
  ) {}

  async getProducts() {
    return await this.productModel.find();
  }
  async getSpecificProducts(id: string) {
    return await this.productModel.findById(id);
  }

  async createProducts(product: Product) {
    const newProduct = new this.productModel(product);
    return await newProduct.save();
  }

  async deleteProducts(id: string) {
    return await this.productModel.findByIdAndDelete(id);
  }

  async updateProducts(id: string, newProduct: Product) {
    return await this.productModel.findByIdAndUpdate(id, newProduct, {
      new: true,
      returnOriginal: false,
    });
  }
}
