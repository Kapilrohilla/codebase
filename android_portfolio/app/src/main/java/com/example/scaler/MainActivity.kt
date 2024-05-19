package com.example.scaler

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Button

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        println("Creating Main Activity")

        val skillBtn = findViewById<Button>(R.id.button)

        skillBtn.setOnClickListener(){
            intent = Intent(this, SkillActivity::class.java)

            startActivity(intent)
        }

    }
}