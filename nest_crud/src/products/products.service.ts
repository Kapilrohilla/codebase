import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { Product } from './interface/products.interface';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

@Injectable()
export class ProductService {
  constructor(
    @InjectModel('products') private readonly productModel: Model<Product>,
  ) {}

  async getProducts() {
    try {
      return await this.productModel.find();
    } catch (err) {
      console.log('ERROR: ' + err.message);
      throw new HttpException(
        'INTERNAL SERVER ERROR',
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }
  async getSpecificProducts(id: string) {
    try {
      return await this.productModel.findById(id);
    } catch (err) {
      console.log('ERROR: ' + err.message);
      if (err.name === 'CastError') {
        throw new HttpException(
          'BAD REQUEST - PRODUCT NOT FOUND',
          HttpStatus.BAD_REQUEST,
        );
      } else {
        throw new Error('Unhandled ERROR: ' + err.message);
      }
    }
  }

  async createProducts(product: Product) {
    try {
      const newProduct = new this.productModel(product);
      return await newProduct.save();
    } catch (err) {
      if (err.name === 'ValidationError') {
        console.log('ERROR: ' + err.message);
        throw new HttpException(
          `${err.name} - ${err.message}`,
          HttpStatus.BAD_REQUEST,
        );
      } else {
        throw new Error('Unhandled ERROR: ' + err.message);
      }
    }
  }

  async deleteProducts(id: string) {
    const r = await this.productModel.findByIdAndDelete(id);
    if (r === null) {
      throw new HttpException('Product not found', HttpStatus.BAD_REQUEST);
    }
    return r;
  }

  async updateProducts(id: string, newProduct: Product) {
    try {
      return await this.productModel.findByIdAndUpdate(id, newProduct, {
        new: true,
        returnOriginal: false,
      });
    } catch (err) {
      console.log('ERROR: ' + err.message);
      if (err.name === 'CastError') {
        throw new HttpException(
          'BAD REQUEST - PRODUCT NOT FOUND',
          HttpStatus.BAD_REQUEST,
        );
      } else {
        throw new Error('Unhandled ERROR: ' + err.message);
      }
    }
  }
}
