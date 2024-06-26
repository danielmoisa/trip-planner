
'use client'
import React, { useState } from 'react';
import { Search, Building2, Home, Waves, Mountain } from 'lucide-react';
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

const SearchBar = () => {
  const [location, setLocation] = useState('');

  const handleSearch = () => {
    console.log('Searching for:', location);
   
  };

  return (
    <div className="w-full mx-auto py-32 rounded-xl bg-gradient-to-r from-blue-100 via-green-100 to-purple-100 flex justify-center">
      <div className='max-w-4xl'>
      <h1 className="text-7xl font-bold mb-2 text-center">
        Explore the <span className="text-green-500">diff</span><span className="text-blue-500">ere</span><span className="text-purple-500">nces</span>
      </h1>
      <p className="text-xl text-center mb-8">Travel freely, anywhere you want.</p>
      
      <div className="bg-white rounded-full shadow-lg p-2 flex items-center ">
        <div className="relative flex-grow">
          <Search className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" />
          <Input
            type="text"
            placeholder="Where to go?"
            className="pl-12 pr-4 py-3 w-full rounded-full border-none"
            value={location}
            onChange={(e) => setLocation(e.target.value)}
          />
        </div>
        <Button 
          className="ml-2 px-6 py-3 rounded-full bg-blue-500 hover:bg-blue-600 text-white"
          onClick={handleSearch}
        >
          Search
        </Button>
      </div>
      
      <div className="flex justify-center mt-4 space-x-8">
        <div className="flex flex-col items-center">
          <Building2 className="w-6 h-6 text-gray-600" />
          <span className="text-sm mt-1">City</span>
        </div>
        <div className="flex flex-col items-center">
          <Home className="w-6 h-6 text-gray-600" />
          <span className="text-sm mt-1">Village</span>
        </div>
        <div className="flex flex-col items-center">
          <Waves className="w-6 h-6 text-gray-600" />
          <span className="text-sm mt-1">Sea</span>
        </div>
        <div className="flex flex-col items-center">
          <Mountain className="w-6 h-6 text-gray-600" />
          <span className="text-sm mt-1">Mountain</span>
        </div>
      </div>
      </div>
    </div>
  );
};

export default SearchBar;